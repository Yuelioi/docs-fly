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
import img10 from '@/assets/images/divergence/10.jpg'

export function getImages() {
    const imgList = [img0, img1, img2, img3, img4, img5, img6, img7, img8, img9, img10]
    const imageCache: { [key: string]: string } = {}
    for (let i = 0; i < 11; i++) {
        if (i === 10) {
            imageCache['.'] = imgList[i]
        } else {
            imageCache[i] = imgList[i]
        }
    }
    return imageCache
}
