<script lang="ts">
    import { Events } from "@wailsio/runtime";
    import { locales, localization } from "../../managers/localization";
    import {
        loadingSuccess,
        news,
        readingNews,
        getNew,
        GetStyle,
        openNewsFeed,
        selectedNewId,
    } from "../../managers/news";
    import { stripMarkup, formatNewsMarkdownToHTML } from "../../utils/string";
    import { Icons } from "../../utils/icons";
    import { Classes, Event, IDs } from "../../utils/consts";
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
            {#await getNew($selectedNewId) then item}
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

                        <div
                            class="{Classes.Views.News.popupContent} {Classes
                                .Global.text}"
                        >
                            {@html formatNewsMarkdownToHTML(item.Content)}
                        </div>
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
