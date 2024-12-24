import { fail } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions = {
    default: async ({ request }) => {
        console.log(request)
        const data = await request.formData();
        const source= data.get('source');
        const destination= data.get('destination');

        if (!source) {
            return fail(400, {source, missing: true });
        }

        if (!destination) {
            return fail(400, {destination, missing: true})
        }

        return { success: true };
    },
} satisfies Actions;