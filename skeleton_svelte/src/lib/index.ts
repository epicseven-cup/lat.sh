// place files you want to import through the `$lib` alias in this folder.
export function UserCreateUrlValidation(s:string):Boolean {
    // const re = new RegExp("\/|\\|\.{2,}|\s")
    const re = new RegExp("^[a-zA-Z]+$")
    return re.test(s)
}
