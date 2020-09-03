package com.lightstep.examples.server;

import org.eclipse.jetty.server.Server;
import org.eclipse.jetty.server.Handler;
import org.eclipse.jetty.server.handler.ContextHandlerCollection;

public class App
{
  public static void main( String[] args )
      throws Exception
    {
      ContextHandlerCollection handlers = new ContextHandlerCollection();
      handlers.setHandlers(new Handler[] {
        new ApiContextHandler(),
      });
      Server server = new Server(8083);
      server.setHandler(handlers);

      server.start();
      server.dumpStdErr();
      server.join();
    }
}
