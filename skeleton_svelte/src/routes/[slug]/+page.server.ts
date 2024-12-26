import {redirect} from "@sveltejs/kit";

export function load({ params }) {
    console.log(params.slug)
    redirect(301, `api/redirect/${params.slug}`)
}