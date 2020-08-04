package com.lightstep.examples.client;

import java.io.IOException;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;

public class App
{
  public static void main( String[] args )
      throws Exception
    {
      String targetUrl = System.getenv("TARGET_URL");
      if (targetUrl == null || targetUrl.length() == 0)
        targetUrl = "http://127.0.0.1:8083";

      OkHttpClient client = new OkHttpClient();
      Request req = new Request.Builder()
        .url(targetUrl + "/content")
        .build();

      while (true) {
        try (Response res = client.newCall(req).execute()) {
          String retval = res.body().string();
          System.out.println(String.format("Request to %s, got %s bytes",
                targetUrl, retval.length()));
        } catch (Exception e) {
          System.out.println(String.format("Request to %s failed: %s",
                targetUrl, e));
        }

        Thread.sleep(1000);
      }
    }
}
