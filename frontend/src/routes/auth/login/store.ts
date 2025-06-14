import { storable } from '$lib/storable';
import { writable, type Writable } from 'svelte/store'

const access_token = writable("");
const refresh_token = writable("");


const { set: setAccess, update: updateAccess } = storable("access_token", access_token);
const { set: setRefresh, update: updateRefresh } = storable("refresh_token", refresh_token);

export {
	access_token,
	refresh_token,
	setAccess,
	updateAccess,
	setRefresh,
	updateRefresh,
}
