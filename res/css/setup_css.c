#include <gtk/gtk.h>
void setup_css() {
  GtkCssProvider *provider;
  GdkDisplay *display;
  GdkScreen *screen;
  provider = gtk_css_provider_new();
  display = gdk_display_get_default();
  screen = gdk_display_get_default_screen(display);
  gtk_style_context_add_provider_for_screen(screen,
      GTK_STYLE_PROVIDER(provider),
      GTK_STYLE_PROVIDER_PRIORITY_APPLICATION);
  gtk_css_provider_load_from_data(GTK_CSS_PROVIDER(provider),
      "#ChatBox{ \n"
      "    background-color: #C0FFC0; \n"
      "}\n"
      "\n"
      "#MessageFrame {\n"
      "    background-color: white;\n"
      "}\n"
      ,-1, NULL);
  g_object_unref(provider);
}
