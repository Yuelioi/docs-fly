export function useLocalStorage(item: any, value: any) {
    console.log(typeof item)

    switch (typeof item) {
        case 'string':
            console.log(1)
    }

    // watch(item, () => {
    //     localStorage.setItem(name, JSON.stringify(item.value))
    // })

    return item
}
