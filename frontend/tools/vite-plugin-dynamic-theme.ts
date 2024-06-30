import { Plugin } from 'vite'
import fs from 'fs'
import path from 'path'

export default function generateThemeConfigPlugin(): Plugin {
    return {
        name: 'generate-theme-config-plugin',
        buildStart() {
            const themeDir = path.join(process.cwd(), 'public/themes')
            const themes: { [key: string]: string[] } = fs
                .readdirSync(themeDir, { withFileTypes: true })
                .filter((dirent) => dirent.isDirectory())
                .reduce((acc: { [key: string]: string[] }, dirent) => {
                    const themeName = dirent.name
                    const themePath = path.join(themeDir, themeName)
                    const cssFiles = fs
                        .readdirSync(themePath)
                        .filter((file) => file.endsWith('.css'))
                    acc[themeName] = cssFiles
                    return acc
                }, {})

            const config = {
                themes
            }

            fs.writeFileSync(
                path.join(process.cwd(), 'public/configs/config.json'),
                JSON.stringify(config, null, 2)
            )
        }
    }
}
