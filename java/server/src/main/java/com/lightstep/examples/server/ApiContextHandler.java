package com.lightstep.examples.server;

import java.io.IOException;
import java.io.PrintWriter;
import java.util.Random;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import org.eclipse.jetty.servlet.ServletContextHandler;
import org.eclipse.jetty.servlet.ServletHolder;

public class ApiContextHandler extends ServletContextHandler
{
  public ApiContextHandler()
  {
    addServlet(new ServletHolder(new ApiServlet()), "/content");
  }

  static final class ApiServlet extends HttpServlet
  {
    static final String LETTERS = "abcdefghijklmnopqrstuvwxyz";
    final Random rand = new Random();

    @Override
    public void doGet(HttpServletRequest req, HttpServletResponse res)
      throws ServletException, IOException
    {
      try (PrintWriter writer = res.getWriter()) {
        writer.write(createRandomString());
      }
    }

    String createRandomString() {
      int length = rand.nextInt(1023) + 1;
      StringBuilder sb = new StringBuilder(length);

      for (int i = 0; i < length; i++) {
        sb.append(LETTERS.charAt(rand.nextInt(LETTERS.length())));
      }

      return sb.toString();
    }
  }
}
