import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ cookies }) => {
	cookies.delete('access_token', { path: '/' });
	return { success: true };
};
