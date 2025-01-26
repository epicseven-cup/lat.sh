// place files you want to import through the `$lib` alias in this folder.
export function UserCreateUrlValidation(s:FormDataEntryValue):Boolean {
    // Default to False when given a file type
    if (s instanceof File){
        return false
    }

    const re = new RegExp("^[a-zA-Z0-9_]+$")
    return re.test(s)
}
