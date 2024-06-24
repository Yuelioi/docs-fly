// https://github.com/dogancelik/divergence-meter

import img0 from '@/assets/images/divergence/0.jpg'
import img1 from '@/assets/images/divergence/1.jpg'
import img2 from '@/assets/images/divergence/2.jpg'
import img3 from '@/assets/images/divergence/3.jpg'
import img4 from '@/assets/images/divergence/4.jpg'
import img5 from '@/assets/images/divergence/5.jpg'
import img6 from '@/assets/images/divergence/6.jpg'
import img7 from '@/assets/images/divergence/7.jpg'
import img8 from '@/assets/images/divergence/8.jpg'
import img9 from '@/assets/images/divergence/9.jpg'
import img11 from '@/assets/images/divergence/11.jpg'

const imagePaths = [img0, img1, img2, img3, img4, img5, img6, img7, img8, img9, img11]

export function getImageUrl(n: string) {
    if (n === '.') {
        return `${imagePaths[10]}`
    } else {
        return `${imagePaths[parseInt(n)]}`
    }
}
