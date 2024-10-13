<script lang="ts">
  import { SaveConfig, LoadState, LoadData, DeployVersion, GetConfig } from '../wailsjs/go/main/App.js';
  import { main } from '../wailsjs/go/models';
  import ReleaseList from './ReleaseLists.svelte';
  import ReleaseHistoryGraph from './ReleaseHistoryGraph.svelte';
  import SettingsModal from './SettingsModal.svelte';
  import SettingsButton from './SettingsButton.svelte';
  import { onMount } from 'svelte';

  let config: main.Config = {
    github_token: '',
    owner: '',
    repo: '',
    table_name: '',
  };
  let releases: main.Release[] = [];
  let version: string = '';
  let showSettingsModal = false;

  onMount(run);

  async function run() {
    const hasMissingConfig = await LoadState();
    config = await GetConfig();
    if (!hasMissingConfig) {
      await loadReleasesAndVersion();
    } else {
      showSettingsModal = true;
    }
  }

  async function loadReleasesAndVersion() {
    const result = await LoadData();
    releases = result.Releases;
    version = result.Version;
  }

  async function deploy(version: string) {
    await DeployVersion(version);
    await loadReleasesAndVersion();
  }

  function openSettings() {
    showSettingsModal = true;
  }

  function closeSettings() {
    showSettingsModal = false;
  }

  async function saveSettings(event: CustomEvent<main.Config>) {
    await SaveConfig(event.detail);
    closeSettings();
    run();
  }
</script>

<main>
  <header class="header">
    <SettingsButton onClick={openSettings} />
    {#if version}
      <p class="version">Online Version: <span class="highlight">{version}</span></p>
    {/if}
  </header>
  {#if showSettingsModal}
    <SettingsModal
      {config}
      on:save={saveSettings}
      on:close={closeSettings}
    />
  {/if}

  <div class="container">
    <ReleaseList {releases} onDeploy={deploy} />
    <h2 class="section-title">Release Graph</h2>
    <ReleaseHistoryGraph {releases} />
  </div>
</main>

<style lang="scss">
  @import url('https://fonts.googleapis.com/css2?family=Fira+Code:wght@400;500;700&display=swap');

  main {
    position: relative;
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 726px;
    padding: 10px;
    height: 500px;
    background-color: transparent;
    font-family: 'Fira Code', monospace;
  }

  .header {
    width: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 20px;
    box-sizing: border-box;
  }

  .version {
    margin: 0;
    font-size: 1.2rem; /* Increased font size */
    color: #ccc; /* Darker color */
    font-weight: bold; /* Bold text */
    .highlight {
      color: #1890ff;
      font-size: 1.4rem; /* Larger font size for the version number */
      font-weight: bold; /* Bold text */
    }
  }

  .container {
    width: 100%;
    box-sizing: border-box;
    background: transparent;
    flex: 1;
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  .section-title {
    margin: 20px 0 10px;
    font-size: 1.5rem;
    color: #ccc;
    text-align: left;
    width: 100%;
    font-weight: bold;
    font-family: 'Fira Code', monospace;
  }
</style>