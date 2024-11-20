<script lang="ts">
	import './../app.css';

	let serverStatus = '-';
	async function checkServer() {
		console.log('Checking your golang server...');

		try {
			const serverUrl = 'http://localhost:8000/api',
				res = await fetch(serverUrl);
			serverStatus = await res.text();
		} catch (e) {
            serverStatus = "Error"
            console.error(e)
		} finally {
            
		}
	}
</script>

<div class="w-100 h-screen flex items-center justify-center flex-col">
	<p>Server Status : {serverStatus}</p>
	<button on:click={checkServer} class="bg-blue-500 rounded-md p-2 text-white mt-8">
		Hit Me to Check server
	</button>
</div>
