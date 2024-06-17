import { Plugin } from 'vite'
import fs from 'fs'
import path from 'path'

export default function generateThemeConfigPlugin(): Plugin {
    return {
        name: 'generate-theme-config-plugin',
        buildStart() {
            const themeDir = path.join(process.cwd(), 'src/theme')
            const themes = fs
                .readdirSync(themeDir, { withFileTypes: true })
                .filter((dirent) => dirent.isDirectory())
                .map((dirent) => dirent.name)

            const config = {
                themes
            }

            fs.writeFileSync(
                path.join(process.cwd(), 'public/configs/themeConfig.json'),
                JSON.stringify(config, null, 2)
            )
        }
    }
}
