import {redirect} from "@sveltejs/kit";

export function load({ params, url }) {
	//console.log(params)
	//console.log(url)
	let o = url.searchParams.get('o')
	let overview = url.searchParams.get('overview')
	if (o != null || overview != null){
		redirect(301, `api/redirect/${params.slug}?o=true`)
	} else {
		redirect(301, `api/redirect/${params.slug}`)
	}
}
