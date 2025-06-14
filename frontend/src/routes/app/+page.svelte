<script>
	import AdminDashboard from "./admin-dashboard.svelte";

	let { data } = $props();
</script>

<h1>Hello {data.access_token ? "friend" : "stranger"}!</h1>
<AdminDashboard />
