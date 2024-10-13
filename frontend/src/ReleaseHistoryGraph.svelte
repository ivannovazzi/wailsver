<script lang="ts">
  import { main } from '../wailsjs/go/models';
  import { onMount, afterUpdate } from 'svelte';
  import Tooltip from './Tooltip.svelte';

  export let releases: main.Release[] = [];

  let graphData = [];
  let weeks = [];

  function prepareData() {
    const today = new Date();
    const oneYearAgo = new Date(today);
    oneYearAgo.setFullYear(today.getFullYear() - 1);

    
    // Adjust end date to the next Saturday to include the whole current week
    let endDate = new Date(today);
    endDate.setDate(endDate.getDate() + (6 - endDate.getDay()));

    // Adjust start date to be 365 days before the end date
    let startDate = new Date(endDate);
    startDate.setDate(endDate.getDate() - 363);

    
    // Map release dates to counts
    let releaseDates = releases.reduce((acc, release) => {
      const dateStr = release.created_at.substring(0, 10); // Get YYYY-MM-DD
      acc[dateStr] = (acc[dateStr] || 0) + 1;
      return acc;
    }, {});

    // Generate list of dates from startDate to endDate
    let date = new Date(startDate);
    let datesArray = [];

    while (date <= endDate) {
      const dateStr = date.toISOString().substring(0, 10);
      datesArray.push({
        date: dateStr,
        count: releaseDates[dateStr] || 0,
        isFuture: date > today,
      });
      date.setDate(date.getDate() + 1);
    }
    
    graphData = datesArray;

    // Organize data into weeks starting on Sunday
    weeks = [];
    let week = [];
    for (let i = 0; i < graphData.length; i++) {
      week.push(graphData[i]);
      if (week.length === 7) {
        weeks.push(week);
        week = [];
      }
    }
    // Add the last week if it's not full
    if (week.length > 0) {
      weeks.push(week);
    }

    console.log(weeks);
  }



  onMount(() => {
    prepareData();
  });

  afterUpdate(() => {
    prepareData();
  });

  // Helper function to get color based on count
  function getColor(count) {
    if (count === 0) return '#ebedf0'; // Lightest color
    else if (count <= 1) return '#c6e48b';
    else if (count <= 3) return '#7bc96f';
    else if (count <= 5) return '#239a3b';
    else return '#196127'; // Darkest color
  }
</script>

<div class="history-graph">
  {#each weeks as week}
    <div class="week">
      {#each week as day}
        <Tooltip text="{day.date}: {day.count} release{day.count !== 1 ? 's' : ''}" position="left">
          <div
            class="day"
            style="
              background-color: {day.isFuture ? '#fff' : getColor(day.count)};
              opacity: {day.count === 0 ? 0.2 : 1};
            "
          ></div>
        </Tooltip>
      {/each}
    </div>
  {/each}
</div>

<style>
  .history-graph {
    display: flex;
    flex-direction: row;
    width: 100%;
  }

  .week {
    display: flex;
    flex-direction: column;
  }

  .day {
    width: 12px;
    height: 12px;
    margin: 1px;
    box-sizing: border-box;
    border-radius: 2px;
  }

  .day:hover {
    outline: 2px solid #000;
  }
</style>