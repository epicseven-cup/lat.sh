import type {Actions} from './$types';
import type {RequestHandler} from "@sveltejs/kit";
import {json} from "node:stream/consumers";

class Result {
    message: string;
    success: boolean;

    constructor(message: string, success: boolean) {
        this.message = message
        this.success = success
    }
}

// export const actions = {
//     default: async ({request}) => {
//         console.log(request)
//         const formData = await request.formData();
//         const source = formData.get('source');
//         const destination = formData.get('destination');
//
//         if (!source) {
//             return new Result("Source is missing", false)
//         }
//
//         if (!destination) {
//             return new Result("Destination is missing", false)
//         }
//
//         const data = {
//             source: source,
//             destination: destination
//         }
//
//         let respond: Response = await fetch("./api/redirect/generate", {
//             method: "POST",
//             body: JSON.stringify(data)
//         })
//
//         let message: string = respond.body == null ? "Something went wrong" : await respond.text()
//
//         return new Result(message, respond.status == 200)
//     },
// } satisfies Actions;