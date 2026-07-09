<script lang="ts">
    import { onMount } from "svelte";
    import { Events } from "@wailsio/runtime";
    import {
        initLocalization,
        locales,
        localization,
    } from "../../managers/localization";
    import { stripMarkup } from "../../utils/string";
    import { base64ToBlob } from "../../utils/helper";
    import { Icons } from "../../utils/icons";
    import {
        GetAllNews,
        GetLatestNewDate,
        GetNewsImage,
    } from "../../../bindings/seegolauncher/internal/services/cacheservice";

    interface NewsItem {
        Title: string;
        Date: string;
        Content: string;
        ImageName: string;
        Image: string;
    }

    let news: NewsItem[] = [];
    let latestNewDate = "";
    let loadingSuccess: boolean = true;

    onMount(async () => {
        initLocalization();
        setNews();
    });

    async function setNews() {
        try {
            let allNews = (await GetAllNews()) as NewsItem[];

            for (const item of allNews) {
                const imageData = await GetNewsImage(item.ImageName);

                const blob = base64ToBlob(imageData);
                item.Image = blob;
            }
            news = allNews;

            latestNewDate = await GetLatestNewDate();
            if (allNews.length === 0 || latestNewDate == "") {
                loadingSuccess = false;
            }
        } catch (err) {
            Events.Emit("feedback", `Failed to load news: ${err}`);
        }
    }

    const style = ["big", "medium", "medium"];
    function GetStyle(i: number) {
        return style[i % style.length];
    }
</script>

<main>
    {#if loadingSuccess}
        <div id="news-view">
            <div class="news-layout">
                {#each news as newItem, i}
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
    {:else}
        <div class="error-view">
            {@html Icons.UI.Alert}
            <p class="error-title text">
                {$locales[localization.NewsLoadFailed]}
            </p>
            <p class="error-comment comment">
                {$locales[localization.NewsLoadFailedDesc]}
            </p>
            <button class="button interactive" on:click={() => setNews()}>
                {$locales[localization.Retry]}
            </button>
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
