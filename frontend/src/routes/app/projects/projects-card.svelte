<script lang="ts">
	import {
		Card,
		CardHeader,
		CardTitle,
		CardDescription,
		CardContent,
	} from "$lib/components/ui/card";
	import { Badge } from "$lib/components/ui/badge";
	import { Button } from "$lib/components/ui/button";
	import { format } from "date-fns";
	import type { Project } from "$lib/types/projects";
	import ProjectViewButton from "./project-view-button.svelte";

	export let project: Project;
</script>

<Card>
	<CardHeader>
		<CardTitle class="text-lg">{project.name}</CardTitle>
		<CardDescription>{project.description}</CardDescription>
	</CardHeader>
	<CardContent>
		<div class="space-y-2 text-sm">
			<p><span class="font-medium">Owner:</span> {project.ownerID}</p>
			<p><span class="font-medium">Tasks:</span> {project.task_count}</p>
			<p>
				<span class="font-medium">Open Tasks:</span>
				<Badge variant={project.open_task_count > 0 ? "default" : "secondary"}>
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
		<ProjectViewButton {project} className="mt-4 w-full" />
	</CardContent>
</Card>
