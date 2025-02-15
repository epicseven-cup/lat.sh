import {redirect} from "@sveltejs/kit";

export function load({ params, url }) {
	//console.log(params)
	//console.log(url)
	let o = url.searchParams('o')
	let overview = url.searchParams('overview')
	if (o != null or overview != null){
		redirect(301, `api/redirect/${params.slug}?o=true`)
	} else {
		redirect(301, `api/redirect/${params.slug}`)
	}
}
