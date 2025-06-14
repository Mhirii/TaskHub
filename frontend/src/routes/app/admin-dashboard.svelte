<script lang="ts">
	import { goto } from "$app/navigation";
	import { Button } from "$lib/components/ui/button";
	import {
		Card,
		CardContent,
		CardHeader,
		CardTitle,
	} from "$lib/components/ui/card";
	import {
		Table,
		TableBody,
		TableCell,
		TableHead,
		TableHeader,
		TableRow,
	} from "$lib/components/ui/table";

	// Sample data for projects and activity log
	const projects = [
		{ id: 1, name: "Project A", status: "Active", tasks: 25, progress: "80%" },
		{
			id: 2,
			name: "Project B",
			status: "In Progress",
			tasks: 15,
			progress: "60%",
		},
		{
			id: 3,
			name: "Project C",
			status: "Completed",
			tasks: 40,
			progress: "100%",
		},
	];

	const activityLog = [
		{
			id: 1,
			user: "Terry Franci",
			action: "Updated Project A",
			time: "5 min ago",
		},
		{
			id: 2,
			user: "Alena Franci",
			action: "Created Task #123",
			time: "10 min ago",
		},
		{
			id: 3,
			user: "Jocelyn Kenter",
			action: "Completed Project C",
			time: "1 hr ago",
		},
	];

	// Navigation function
	function goToHome() {
		goto("/home");
	}
</script>

<div class="max-w-7xl mx-auto p-6 space-y-8">
	<!-- Header -->
	<div class="flex justify-between items-center">
		<h1 class="text-3xl font-bold">TaskHub Admin Dashboard</h1>
		<Button onclick={goToHome}>Go to Home</Button>
	</div>

	<!-- Projects Section -->
	<Card>
		<CardHeader>
			<CardTitle>Projects</CardTitle>
		</CardHeader>
		<CardContent>
			<Table>
				<TableHeader>
					<TableRow>
						<TableHead>Name</TableHead>
						<TableHead>Status</TableHead>
						<TableHead>Tasks</TableHead>
						<TableHead>Progress</TableHead>
					</TableRow>
				</TableHeader>
				<TableBody>
					{#each projects as project}
						<TableRow>
							<TableCell>{project.name}</TableCell>
							<TableCell>{project.status}</TableCell>
							<TableCell>{project.tasks}</TableCell>
							<TableCell>{project.progress}</TableCell>
						</TableRow>
					{/each}
				</TableBody>
			</Table>
		</CardContent>
	</Card>

	<!-- Activity Log Section -->
	<Card>
		<CardHeader>
			<CardTitle>Activity Log</CardTitle>
		</CardHeader>
		<CardContent>
			<ul class="space-y-4">
				{#each activityLog as log}
					<li class="text-sm">
						<span class="font-semibold">{log.user}</span>
						{log.action} <span class="text-muted-foreground">({log.time})</span>
					</li>
				{/each}
			</ul>
		</CardContent>
	</Card>
</div>
