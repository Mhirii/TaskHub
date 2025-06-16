import type { PageServerLoad } from "./$types";
import { redirect } from '@sveltejs/kit'


export const load: PageServerLoad = async ({ cookies }) => {
	if (cookies.get('access_token')) {
		redirect(302, '/app');
	} else {
		redirect(302, '/auth/login');
	}
};
