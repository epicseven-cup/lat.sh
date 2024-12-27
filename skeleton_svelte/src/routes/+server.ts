// import type {RequestHandler} from "@sveltejs/kit";
// import { json } from '@sveltejs/kit';
// export const POST: RequestHandler = async (event) => {
//     const body:FormData = await event.request.formData()
//
//     let source = body.get('source')
//     if (source == null){
//         return json({status: false, message: "Source is empty"})
//     }
//
//     let destination = body.get('destination')
//     if (destination == null){
//         return json({status: false, message: "Destination is empty"})
//     }
//
//
//     let data = {
//         source: source,
//         destination: destination
//     }
//
//
//     let respond: Response = await fetch("/api/generate", {
//         method: "POST",
//         body: JSON.stringify(data)
//     })
//
//     return json({status: respond.ok, message: respond.body})
// }
//
