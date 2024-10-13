<script lang="ts">
  import { main } from '../wailsjs/go/models';
  import { format, formatDistanceToNow } from 'date-fns';
  import { BrowserOpenURL } from '../wailsjs/runtime/runtime';

  export let releases: main.Release[] = [];
  export let onDeploy: (version: string) => void;

  let loadingVersion: string | null = null;

  function formatDate(dateString: string): string {
    return format(new Date(dateString), 'yyyy-MM-dd HH:mm');
  }

  function timeAgo(dateString: string): string {
    return formatDistanceToNow(new Date(dateString), { addSuffix: true });
  }

  function openInBrowser(url: string) {
    BrowserOpenURL(url);
  }

  function extractVersion(tagName: string): string {
    return tagName.split('v')[1];
  }

  async function deployVersion(release: main.Release) {
    let version = extractVersion(release.tag_name);
    loadingVersion = version;
    await onDeploy(version);
    loadingVersion = null;
  }

  function isLoading(release): boolean {
    return loadingVersion === extractVersion(release.tag_name);
  }
</script>

<div class="releases">
  {#each releases as release, index}
    <div class="release">
      <span class="release-star">
        {#if release.Current }
          <!-- SVG Star Icon -->
          <svg width="16" height="16" viewBox="0 0 24 24" fill="#FFD700" xmlns="http://www.w3.org/2000/svg">
            <path d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z"/>
          </svg>
        {/if}
      </span>
      <span class="release-name">{release.name}</span>
      <span class="release-info">
        {formatDate(release.created_at)}
        <span class="time-ago">({timeAgo(release.created_at)})</span>
      </span>
      <span class="release-tag"><code>{release.tag_name}</code></span>
      <a
        href="#"
        on:click={() => openInBrowser(release.html_url)}
        class="release-link"
      >
        View Release
      </a>
      <div class="release-deploy">
        {#if !release.Current}
          <button class="deploy-button" on:click={() => deployVersion(release)} disabled={isLoading(release)}>
            {#if isLoading(release)}
              Deploying...
            {:else}
              Deploy
            {/if}
          </button>
        {/if}
      </div>
    </div>
  {/each}
</div>

<style lang="scss">
  @import url('https://fonts.googleapis.com/css2?family=Fira+Code:wght@400;500;700&display=swap');

  .releases {
    display: flex;
    flex-direction: column;
    max-height: 360px;
    overflow-y: auto;
    width: 100%; /* Added width */
    &::-webkit-scrollbar {
      display: none;
    }

    -ms-overflow-style: none; /* IE and Edge */
    scrollbar-width: none; /* Firefox */
  }

  .release {
    background-color: rgba(51, 51, 51, 0.3);
    padding: 5px;
    font-size: 0.85rem; /* Reduced font size */
    display: grid;
    grid-template-columns: 24px 2fr 3fr 2fr auto 80px;
    margin-bottom: 4px;
    align-items: center;
    position: relative;
    gap: 10px; /* Add spacing between columns */
    height: 50px; /* Fixed height for rows */

    &:hover .release-link,
    &:hover .deploy-button {
      opacity: 1;
    }

    .release-star {
      width: 16px;
      display: flex;
      align-items: center;
      justify-content: center;
    }

    .release-name {
      font-size: 0.95rem; /* Reduced font size */
      font-weight: 500;
      color: #1890ff;
      text-align: left;
    }

    .release-info,
    .release-tag,
    .release-link,
    .deploy-button {
      font-size: 0.85rem; /* Reduced font size */
      color: #ccc;
    }

    .release-info {
      display: flex;
      flex-direction: column;
    }

    .release-info .time-ago {
      font-size: 0.75rem; /* Smaller font size */
      opacity: 0.6; /* Less visible */
    }

    .release-tag {
      code {
  
        background-color: rgba(85, 85, 85, 0.6);
        padding: 2px 4px;
        color: #fff;
        font-weight: bold;
        border: 1px solid rgba(102, 102, 102, 0.4);
        font-size: 0.75rem; /* Reduced font size */
      }
    }
    .release-deploy {
      display: flex;
      justify-content: flex-end;
    }

    .release-link {
      color: #1890ff;
      text-decoration: none;
      font-weight: 500;
      position: relative;
      right: 10px;
      opacity: 0;
      transition: opacity 0.3s ease-in-out;

      &::after {
        content: '';
        position: absolute;
        width: 100%;
        height: 2px;
        bottom: -2px;
        left: 0;
        background-color: #1890ff;
        visibility: hidden;
        transform: scaleX(0);
        transition: all 0.3s ease-in-out;
      }

      &:hover::after {
        visibility: visible;
        transform: scaleX(1);
      }

      &:hover {
        color: #40a9ff;
      }
    }

    .deploy-button {
      background-color: #333;
      border: 1px solid #555;
      color: #fff;
      font-weight: 500;
      cursor: pointer;
      opacity: 0;
      transition: opacity 0.3s ease-in-out, background-color 0.3s, border-color 0.3s;
      padding: 5px 10px;
      border-radius: 4px;
      margin-left: 10px;

      font-size: 0.85rem; /* Reduced font size */

      &:hover {
        background-color: #444;
        border-color: #666;
      }

      &:disabled {
        cursor: not-allowed;
        background-color: #555;
        border-color: #777;
      }
    }
  }
</style>