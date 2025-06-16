<script lang="ts">
	import DataTable from "./data-table.svelte";
	import { columns } from "./columns";
	import { Button } from "$lib/components/ui/button";
	import { Input } from "$lib/components/ui/input";

	import { Search, Plus } from "@lucide/svelte";
	import CreateUserForm from "./create-user-form.svelte";
	import { Dialog, DialogContent } from "$lib/components/ui/dialog";
	import DialogTrigger from "$lib/components/ui/dialog/dialog-trigger.svelte";
	import DialogHeader from "$lib/components/ui/dialog/dialog-header.svelte";
	import DialogTitle from "$lib/components/ui/dialog/dialog-title.svelte";

	let { data } = $props();
	let searchQuery = "";

	let filteredUsers = $derived(
		data.users.filter((user) =>
			Object.values(user).some(
				(value) =>
					value &&
					value.toString().toLowerCase().includes(searchQuery.toLowerCase()),
			),
		),
	);
</script>

<div class="container mx-auto py-8 space-y-6">
	<div
		class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4"
	>
		<div>
			<h1 class="text-3xl font-bold tracking-tight">Users</h1>
			<p class="text-muted-foreground">Manage users and their permissions.</p>
		</div>
		<Dialog>
			<DialogTrigger>
				<Button>
					<Plus class="w-4 h-4 mr-2" />
					Add User
				</Button>
			</DialogTrigger>
			<DialogContent>
				<DialogHeader>
					<DialogTitle>Create User</DialogTitle>
				</DialogHeader>
				<CreateUserForm {data} />
			</DialogContent>
		</Dialog>
	</div>

	<div class="flex flex-col sm:flex-row gap-4 items-center">
		<div class="relative w-full sm:w-72">
			<Search class="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
			<Input
				type="search"
				placeholder="Search users..."
				class="pl-8"
				bind:value={searchQuery}
			/>
		</div>
		<!-- Add more filters here if needed -->
	</div>

	<div class="rounded-md border">
		<DataTable data={filteredUsers} {columns} />
	</div>
</div>
