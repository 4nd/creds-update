import {defineConfig, Plugin, ResolvedConfig} from "vite";
import * as fs from "node:fs";

function hotFilePlugin(): Plugin[] {
    let exitHandlersBound = false
    let resolvedConfig: ResolvedConfig
    return [{
        name: 'hotFile-on-serve',
        apply: 'serve',
        configResolved(config) {
            resolvedConfig = config
        },
        configureServer(server) {
            let hotFile = resolvedConfig.root + '/vite-hot';
            server.httpServer?.once('listening', () => {
                const address = server.httpServer?.address()
                if (typeof address === 'object') {
                    const config = resolvedConfig
                    const configHmrProtocol = typeof config.server.hmr === 'object' ? config.server.hmr.protocol : null
                    const clientProtocol = configHmrProtocol ? (configHmrProtocol === 'wss' ? 'https' : 'http') : null
                    const serverProtocol = config.server.https ? 'https' : 'http'
                    const protocol = clientProtocol ?? serverProtocol

                    const configHmrHost = typeof config.server.hmr === 'object' ? config.server.hmr.host : null
                    const configHost = typeof config.server.host === 'string' ? config.server.host : null
                    const serverAddress = address.family === '6' || address.family === 'IPv6' ? `[${address.address}]` : address.address
                    const host = configHmrHost ?? configHost ?? serverAddress

                    const configHmrClientPort = typeof config.server.hmr === 'object' ? config.server.hmr.clientPort : null
                    const port = configHmrClientPort ?? address.port

                    fs.writeFileSync(hotFile, `${protocol}://${host}:${port}`)
                    if (!exitHandlersBound) {
                        const clean = () => {
                            if (fs.existsSync(hotFile)) {
                                fs.rmSync(hotFile)
                            }
                        }

                        process.on('exit', clean)
                        process.on('SIGINT', () => process.exit())
                        process.on('SIGTERM', () => process.exit())
                        process.on('SIGHUP', () => process.exit())

                        exitHandlersBound = true
                    }
                }
            })
        }
    }];
}

export default defineConfig({
    build: {
        manifest: true,
        rollupOptions: {
            input: ['assets/src/app.js'],
        },
        copyPublicDir: false,
        outDir: 'assets/dist',
        assetsDir: '',
    },
    plugins: [
        hotFilePlugin(),
    ]
})