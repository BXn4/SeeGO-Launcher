<script lang="ts">
    import { onMount } from "svelte";
    import { Events } from "@wailsio/runtime";
    import {
        initLocalization,
        locales,
        localization,
    } from "../../managers/localization";
    import { loadingSuccess, news } from "../../managers/news";
    import { stripMarkup } from "../../utils/string";
    import { Icons } from "../../utils/icons";
    import { Event, View } from "../../utils/consts";

    let latestNewDate = "";

    onMount(() => {
        initLocalization();

        const readLatest = Events.On(Event.Main.News.readLatest, async (e) => {
            alert("a");
        });

        return () => {
            readLatest();
        };
    });

    const style = ["big", "medium", "medium"];
    function GetStyle(i: number) {
        return style[i % style.length];
    }
</script>

<main>
    {#if !loadingSuccess}
        <div class="error-view">
            {@html Icons.UI.Alert}
            <p class="error-title text">
                {$locales[localization.NewsLoadFailed]}
            </p>
            <p class="error-comment comment">
                {$locales[localization.NewsLoadFailedDesc]}
            </p>
            <button
                class="button interactive"
                onclick={() => Events.Emit(Event.Global.newsFeedUpdated, null)}
            >
                {$locales[localization.Retry]}
            </button>
        </div>
    {:else}
        <div id="news-view">
            <div class="news-layout">
                {#each $news as newItem, i}
                    <div
                        class="news-card {GetStyle(i)}"
                        style="background-image: url('{newItem.Image}')"
                    >
                        <span class="news-overlay"></span>
                        <div class="news-content">
                            {#if latestNewDate == newItem.Date}
                                <span class="news-badge"
                                    >{$locales[localization.newsLatest]}</span
                                >
                            {:else}
                                <span class="news-badge old"
                                    >{newItem.Date}</span
                                >
                            {/if}
                            <p class="news-title">{newItem.Title}</p>
                            <p class="news-comment">
                                {stripMarkup(newItem.Content)}
                            </p>
                            <button
                                class="button news-read interactive"
                                id="news-read-{i}"
                                >{$locales[localization.newsRead]}</button
                            >
                        </div>
                    </div>
                {/each}
            </div>
        </div>
    {/if}
</main>

<style>
    @import "../../public/styles/views/news.css";
    main {
        width: 100vw;
        height: 100vh;
        display: grid;
        grid-template-columns: var(--width, 240px) 1fr;
        grid-template-rows: var(--height, 48px) 1fr;
        overflow: hidden;
    }
</style>
