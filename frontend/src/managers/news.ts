import {
  GetAllNews,
  GetLatestNew,
  GetLatestNewDate,
  GetNewsImage,
} from "../../bindings/seegolauncher/internal/services/cacheservice";
import { NewsItem } from "../../bindings/seegolauncher/internal/services/models";
import { base64ToBlob } from "../utils/helper";
import { stripMarkup } from "../utils/string";

export async function getAllNews(): Promise<NewsItem[]> {
  const allNews = (await GetAllNews()) as NewsItem[];

  if (allNews.length === 0) {
    return [];
  }

  for (const item of allNews) {
    const imageData = await GetNewsImage(item.ImageName);

    const blob = base64ToBlob(imageData);
    item.Image = blob;

    item.Content = stripMarkup(item.Content)
  }

  return allNews;
}

export async function getLatestNew(): Promise<NewsItem | undefined> {
  const latestNew = (await GetLatestNew()) as NewsItem;

  if (latestNew == undefined) {
    return undefined;
  }

  const imageData = await GetNewsImage(latestNew.ImageName);
  const imageUrl = base64ToBlob(imageData);
  latestNew.Image = imageUrl;

  latestNew.Content = stripMarkup(latestNew.Content)
  return latestNew;
}


export async function getLatestNewDate(): Promise<string> {
  return await GetLatestNewDate();
}
