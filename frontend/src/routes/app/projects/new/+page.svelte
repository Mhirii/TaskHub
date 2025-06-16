<script lang="ts">
	import { type SuperValidated, type Infer } from "sveltekit-superforms";
	import { type ProjectSchema } from "./schema";
	import CreateProjectForm from "./create-project-form.svelte";
	import type { PageData } from "./$types";

	export let data: PageData;
</script>

<CreateProjectForm projectForm={data.form} />
