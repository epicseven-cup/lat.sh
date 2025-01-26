/*
import type { Actions } from './$types';
import {aw} from "vitest/dist/chunks/reporters.D7Jzd9GS";
import {fail} from "@sveltejs/kit";

export const actions = {
    default: async ({request}) => {
        let data: FormData = await request.formData()




        let source: string = data.get("source") ?? ""
        let destination: string = data.get("destination") ?? ""

        if (source === ""){
            return fail(400, { source, missing: true})
        }

        if (destination === ""){
            return fail(400, {destination, missing: true})
        }

    }
} satisfies Actions;*/
