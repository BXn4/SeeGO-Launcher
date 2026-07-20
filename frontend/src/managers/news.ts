import {
  GetAllNews,
  GetNewsImage,
} from "../../bindings/seegolauncher/internal/services/cacheservice";
import { NewsItem } from "../../bindings/seegolauncher/internal/services/models";
import { Event } from "../utils/consts";
import { base64ToBlob } from "../utils/helper";
import { stripMarkup } from "../utils/string";
import { Events } from "@wailsio/runtime";
import { writable } from "svelte/store";

export const news = writable<NewsItem[]>([]);
export const loadingSuccess = writable<boolean | null>(null);

export async function initNews() {
  loadingSuccess.set(await setAllNews())

  Events.On(Event.Global.newsFeedUpdated, async (e) => {
    loadingSuccess.set(await setAllNews())
  })
}

export async function setAllNews() {
  const allNews = (await GetAllNews()) as NewsItem[];

  if (allNews.length === 0) {
    return false;
  }

  for (const item of allNews) {
    const imageData = await GetNewsImage(item.ImageName);

    const blob = base64ToBlob(imageData);
    item.Image = blob;

    item.Content = stripMarkup(item.Content)
  }

  news.set(allNews)
  return true
}
