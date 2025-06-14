<script lang="ts">
	import {
		Table,
		TableHeader,
		TableRow,
		TableHead,
		TableBody,
		TableCell,
	} from "$lib/components/ui/table";
	import { Badge } from "$lib/components/ui/badge";
	import { Button } from "$lib/components/ui/button";
	import { format } from "date-fns";
	import type { Project } from "$lib/types/projects";
	import ProjectViewButton from "./project-view-button.svelte";

	export let projects: Project[];
</script>

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
						variant={project.open_task_count > 0 ? "default" : "secondary"}
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
					<ProjectViewButton {project} size="sm" />
				</TableCell>
			</TableRow>
		{/each}
	</TableBody>
</Table>
