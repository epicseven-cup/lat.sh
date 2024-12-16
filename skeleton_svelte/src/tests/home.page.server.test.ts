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

test('Number Validate User inputted Url', () => {
    let input:string = String.raw`1234567890`
    const result:Boolean = UserCreateUrlValidation(input)
    expect(result).toEqual(true)
})

test('Number Mix Letter Validate User inputted Url', () => {
    let input:string = String.raw`1231adsfasf0u45sdfa67890`
    const result:Boolean = UserCreateUrlValidation(input)
    expect(result).toEqual(true)
})

test('Number Mix Underscore Validate User inputted Url', () => {
    let input:string = String.raw`123_456_7890`
    const result:Boolean = UserCreateUrlValidation(input)
    expect(result).toEqual(true)
})

test('Number, Letter and Underscore Validate User inputted Url', () => {
    let input:string = String.raw`123_4567_890abcd_d_e_4_f`
    const result:Boolean = UserCreateUrlValidation(input)
    expect(result).toEqual(true)
})