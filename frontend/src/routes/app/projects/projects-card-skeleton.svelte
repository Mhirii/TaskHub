<script lang="ts">
	import {
		Card,
		CardContent,
		CardDescription,
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
	import { Badge } from "$lib/components/ui/badge";
	import { Button } from "$lib/components/ui/button";
	import { format } from "date-fns";
	import type { PageData } from "./$types";
	import { Skeleton } from "$lib/components/ui/skeleton";

	export let data: PageData;

	// Skeleton placeholders
	const skeletonRows = Array(5).fill(null);
</script>

<div class="flex justify-center">
	<div class="w-full xl:max-w-[1200px] py-10">
		<Card class="mb-8">
			<CardHeader>
				<CardTitle class="text-3xl font-bold">Projects</CardTitle>
				<CardDescription
					>Manage your projects and track progress</CardDescription
				>
			</CardHeader>
		</Card>

		{#await data.projects}
			<!-- Mobile view: Skeleton Cards -->
			<div class="lg:hidden space-y-4">
				{#each skeletonRows as _}
					<Card>
						<CardHeader>
							<Skeleton class="h-6 w-1/2" />
							<Skeleton class="h-4 w-3/4" />
						</CardHeader>
						<CardContent>
							<div class="space-y-2">
								<Skeleton class="h-4 w-1/3" />
								<Skeleton class="h-4 w-1/4" />
								<Skeleton class="h-4 w-1/4" />
								<Skeleton class="h-4 w-1/3" />
								<Skeleton class="h-4 w-1/3" />
							</div>
							<Skeleton class="h-10 w-full mt-4" />
						</CardContent>
					</Card>
				{/each}
			</div>

			<!-- Desktop view: Skeleton Table -->
			<div class="hidden lg:block">
				<Table>
					<TableHeader>
						<TableRow>
							<TableHead><Skeleton class="h-4 w-20" /></TableHead>
							<TableHead><Skeleton class="h-4 w-32" /></TableHead>
							<TableHead><Skeleton class="h-4 w-20" /></TableHead>
							<TableHead><Skeleton class="h-4 w-16" /></TableHead>
							<TableHead><Skeleton class="h-4 w-24" /></TableHead>
							<TableHead class="hidden xl:table-cell"
								><Skeleton class="h-4 w-20" /></TableHead
							>
							<TableHead class="hidden xl:table-cell"
								><Skeleton class="h-4 w-20" /></TableHead
							>
							<TableHead><Skeleton class="h-4 w-16" /></TableHead>
						</TableRow>
					</TableHeader>
					<TableBody>
						{#each skeletonRows as _}
							<TableRow>
								<TableCell><Skeleton class="h-4 w-24" /></TableCell>
								<TableCell><Skeleton class="h-4 w-40" /></TableCell>
								<TableCell><Skeleton class="h-4 w-20" /></TableCell>
								<TableCell><Skeleton class="h-4 w-12" /></TableCell>
								<TableCell><Skeleton class="h-4 w-16" /></TableCell>
								<TableCell class="hidden xl:table-cell"
									><Skeleton class="h-4 w-20" /></TableCell
								>
								<TableCell class="hidden xl:table-cell"
									><Skeleton class="h-4 w-20" /></TableCell
								>
								<TableCell><Skeleton class="h-8 w-16" /></TableCell>
							</TableRow>
						{/each}
					</TableBody>
				</Table>
			</div>
		{:then projects}
			{#if projects?.length === 0}
				<Card>
					<CardContent class="text-center py-10">
						<p class="text-muted-foreground">
							No projects found. Create your first project!
						</p>
						<Button class="mt-4">Create Project</Button>
					</CardContent>
				</Card>
			{:else}
				<div class="space-y-8">
					<!-- Mobile view: Cards -->
					<div class="lg:hidden space-y-4">
						{#each projects as project}
							<Card>
								<CardHeader>
									<CardTitle class="text-lg">{project.name}</CardTitle>
									<CardDescription>{project.description}</CardDescription>
								</CardHeader>
								<CardContent>
									<div class="space-y-2 text-sm">
										<p>
											<span class="font-medium">Owner:</span>
											{project.ownerID}
										</p>
										<p>
											<span class="font-medium">Tasks:</span>
											{project.task_count}
										</p>
										<p>
											<span class="font-medium">Open Tasks:</span>
											<Badge
												variant={project.open_task_count > 0
													? "default"
													: "secondary"}
											>
												{project.open_task_count}
											</Badge>
										</p>
										<p>
											<span class="font-medium">Created:</span>
											{format(new Date(project.created_at), "MMM dd, yyyy")}
										</p>
										<p>
											<span class="font-medium">Updated:</span>
											{format(new Date(project.updated_at), "MMM dd, yyyy")}
										</p>
									</div>
									<Button class="mt-4 w-full">View Details</Button>
								</CardContent>
							</Card>
						{/each}
					</div>

					<!-- Desktop view: Table -->
					<div class="hidden lg:block w-full">
						<Table>
							<TableHeader>
								<TableRow>
									<TableHead>Name</TableHead>
									<TableHead>Description</TableHead>
									<TableHead>Owner</TableHead>
									<TableHead>Tasks</TableHead>
									<TableHead>Open Tasks</TableHead>
									<TableHead class="hidden xl:table-cell">Created</TableHead>
									<TableHead class="hidden xl:table-cell">Updated</TableHead>
									<TableHead>Actions</TableHead>
								</TableRow>
							</TableHeader>
							<TableBody>
								{#each projects as project}
									<TableRow>
										<TableCell class="font-medium">{project.name}</TableCell>
										<TableCell>{project.description}</TableCell>
										<TableCell>{project.ownerID}</TableCell>
										<TableCell>{project.task_count}</TableCell>
										<TableCell>
											<Badge
												variant={project.open_task_count > 0
													? "default"
													: "secondary"}
											>
												{project.open_task_count}
											</Badge>
										</TableCell>
										<TableCell class="hidden xl:table-cell">
											{format(new Date(project.created_at), "MMM dd, yyyy")}
										</TableCell>
										<TableCell class="hidden xl:table-cell">
											{format(new Date(project.updated_at), "MMM dd, yyyy")}
										</TableCell>
										<TableCell>
											<Button size="sm">View</Button>
										</TableCell>
									</TableRow>
								{/each}
							</TableBody>
						</Table>
					</div>
				</div>
			{/if}
		{:catch error}
			<Card>
				<CardContent class="text-center py-10">
					<p class="text-destructive">
						Error loading projects: {error.message}
					</p>
				</CardContent>
			</Card>
		{/await}
	</div>
</div>
