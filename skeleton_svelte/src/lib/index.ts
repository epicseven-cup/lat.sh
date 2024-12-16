// place files you want to import through the `$lib` alias in this folder.
export function UserCreateUrlValidation(s:string):Boolean {
    const re = new RegExp("^[a-zA-Z0-9_]+$")
    return re.test(s)
}
