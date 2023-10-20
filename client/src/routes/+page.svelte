<script lang="ts">
	import { onMount } from 'svelte';

	type Message = {
		commandName: string;
		messageBody: string;
	};

	let messages = [] as Message[];
	onMount(async () => {
		console.log('connecting');

		const socket = new WebSocket('ws://localhost:1111/ws');

		socket.onmessage = function (event) {
			const message = JSON.parse(event.data);
			messages = [...messages, message];
		};
	});
</script>

<ul>
	{#each messages as message}
		<li>{`[${message.commandName}] ${message.messageBody}`}</li>
	{/each}
</ul>
