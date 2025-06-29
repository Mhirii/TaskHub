
import type { LayoutServerLoad } from "./$types";
import { redirect } from '@sveltejs/kit'


export const load: LayoutServerLoad = async ({ cookies }) => {
	if (cookies.get('access_token')) {
		redirect(302, '/app');
	}
	return {}
};
