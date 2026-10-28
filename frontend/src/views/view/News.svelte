<script lang="ts">
    import { onMount } from "svelte";
    import { Events } from "@wailsio/runtime";
    import { locales, localization } from "../../managers/localization";
    import {
        loadingSuccess,
        news,
        readingNews,
        getNew,
        formatToHTML,
    } from "../../managers/news";
    import { stripMarkup } from "../../utils/string";
    import { Icons } from "../../utils/icons";
    import { Classes, Event, IDs } from "../../utils/consts";

    let selectedNewId: number;

    onMount(() => {
        const readLatest = Events.On(Event.Main.News.readLatest, async (e) => {
            openNewsFeed(0);
        });

        return () => {
            readLatest();
        };
    });

    const style = ["big", "medium", "medium"];
    function GetStyle(i: number) {
        return style[i % style.length];
    }

    function openNewsFeed(id: number) {
        selectedNewId = id;
        readingNews.set(true);
    }

    function closeNewsFeed() {
        readingNews.set(false);
    }
</script>

<main>
    {#if $loadingSuccess == false}
        <div id={IDs.Views.Error.ID}>
            {@html Icons.UI.Alert}
            <p class="{Classes.Views.Error.title} {Classes.Global.text}">
                {$locales[localization.NewsLoadFailed]}
            </p>
            <p class="{Classes.Views.Error.alt} {Classes.Global.comment}">
                {$locales[localization.NewsLoadFailedDesc]}
            </p>
            <button
                class="{Classes.Global.button} {Classes.Global.interactive}"
                onclick={() => Events.Emit(Event.Global.newsFeedUpdated, null)}
            >
                {$locales[localization.Retry]}
            </button>
        </div>
    {:else if $loadingSuccess == true}
        <div id={IDs.Views.News.ID}>
            <div class={Classes.Views.News.layout}>
                {#each $news as item, i}
                    <div
                        class="{Classes.Views.News.card} {GetStyle(i)}"
                        style="background-image: url('{item.Image}')"
                    >
                        <span class={Classes.Views.News.overlay}></span>
                        <div class={Classes.Views.News.content}>
                            {#if i === 0}
                                <span class={Classes.Global.newsBadge}
                                    >{$locales[localization.newsLatest]}</span
                                >
                            {:else}
                                <span
                                    class="{Classes.Global.newsBadge} {Classes
                                        .Global.old}">{item.Date}</span
                                >
                            {/if}
                            <p class={Classes.Global.newsTitle}>{item.Title}</p>
                            <p class={Classes.Global.newsAlt}>
                                {stripMarkup(item.Content)}
                            </p>
                            <button
                                class="{Classes.Global.button} {Classes.Global
                                    .newsRead} {Classes.Global.interactive}"
                                onclick={() => openNewsFeed(i)}
                                >{$locales[localization.newsRead]}</button
                            >
                        </div>
                    </div>
                {/each}
            </div>
        </div>
    {/if}

    {#if $readingNews == true}
        <div class={Classes.Views.News.popupOverlay}>
            {#await getNew(selectedNewId) then item}
                {#if item}
                    <div class={Classes.Views.News.popup}>
                        <div
                            class={Classes.Views.News.popupImage}
                            style="background-image: url('{item.Image}')"
                        >
                            <span class={Classes.Views.News.overlay}></span>
                            <div class={Classes.Views.News.popupHeader}>
                                <span class={Classes.Global.newsBadge}
                                    >{item.Date}</span
                                >
                                <p class={Classes.Global.newsTitle}>
                                    {item.Title}
                                </p>
                            </div>
                        </div>

                        <div class={Classes.Views.News.popupContent}></div>
                    </div>
                {/if}
            {/await}
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
        position: relative;
    }
</style>
