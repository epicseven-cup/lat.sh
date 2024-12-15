import {UserCreateUrlValidation} from "$lib";
import { expect, test } from 'vitest';
test('Slash Validating User inputted Url', () => {
    let input:string = String.raw`localhost:../fakepath`
    const result:Boolean = UserCreateUrlValidation(input)
    expect(result).toEqual(false)
})


test('Single Slash Validating User inputted Url', () => {
    let input:string = String.raw`/`
    const result:Boolean = UserCreateUrlValidation(input)
    expect(result).toEqual(false)
})


test('Dot Validating User inputted Url', () => {
    let input:string = String.raw`google.com`
    const result:Boolean = UserCreateUrlValidation(input)
    expect(result).toEqual(false)
})

// Good testcase
test('Validate User inputted Url', () => {
    let input:string = String.raw`googlecom`
    const result:Boolean = UserCreateUrlValidation(input)
    expect(result).toEqual(true)
})