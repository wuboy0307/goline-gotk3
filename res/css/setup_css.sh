#!/bin/bash


echo 'GEN setup_css.c'
echo '#include <gtk/gtk.h>' > setup_css.c
echo 'void setup_css() {' >> setup_css.c
echo '  GtkCssProvider *provider;' >> setup_css.c
echo '  GdkDisplay *display;' >> setup_css.c
echo '  GdkScreen *screen;' >> setup_css.c
echo '  provider = gtk_css_provider_new();' >> setup_css.c
echo '  display = gdk_display_get_default();' >> setup_css.c
echo '  screen = gdk_display_get_default_screen(display);' >> setup_css.c
echo '  gtk_style_context_add_provider_for_screen(screen,' >> setup_css.c
echo '      GTK_STYLE_PROVIDER(provider),' >> setup_css.c
echo '      GTK_STYLE_PROVIDER_PRIORITY_APPLICATION);'  >> setup_css.c
echo '  gtk_css_provider_load_from_data(GTK_CSS_PROVIDER(provider),' >> setup_css.c
IFS=''
while read line; do
    echo "      \"$line\n\"";
done < theme.css >> setup_css.c
echo '      ,-1, NULL);' >> setup_css.c
echo '  g_object_unref(provider);' >> setup_css.c
echo '}' >> setup_css.c
