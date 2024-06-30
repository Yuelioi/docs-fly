function echo(param?: string): string {
    return param ?? ''
}

interface Product {
    name: string
}

const p: Product = { name: '1' }

function foo(p: Product) {
    console.log(p.name)
}
