package glade
const AutologinWindow string =` <?xml version="1.0" encoding="UTF-8"?> <!-- Generated with glade 3.18.3 --> <interface> <requires lib="gtk+" version="3.12"/> <object class="GtkWindow" id="Window"> <property name="can_focus">False</property> <property name="margin_left">5</property> <property name="title" translatable="yes">Autologin</property> <property name="default_width">300</property> <property name="default_height">200</property> <child> <object class="GtkBox" id="box1"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="margin_left">5</property> <property name="margin_right">5</property> <property name="margin_top">5</property> <property name="margin_bottom">5</property> <property name="orientation">vertical</property> <property name="spacing">10</property> <child> <object class="GtkSpinner" id="spinner1"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="active">True</property> </object> <packing> <property name="expand">True</property> <property name="fill">True</property> <property name="position">0</property> </packing> </child> <child> <object class="GtkLabel" id="label1"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="label" translatable="yes">Login...</property> </object> <packing> <property name="expand">False</property> <property name="fill">True</property> <property name="position">1</property> </packing> </child> <child> <object class="GtkButton" id="ExitButton"> <property name="label" translatable="yes">Exit</property> <property name="visible">True</property> <property name="can_focus">True</property> <property name="receives_default">True</property> </object> <packing> <property name="expand">False</property> <property name="fill">True</property> <property name="position">2</property> </packing> </child> </object> </child> </object> </interface> `
const ChatWindow string =` <?xml version="1.0" encoding="UTF-8"?> <!-- Generated with glade 3.18.3 --> <interface> <requires lib="gtk+" version="3.12"/> <object class="GtkWindow" id="Window"> <property name="can_focus">False</property> <property name="default_width">400</property> <property name="default_height">500</property> <child> <object class="GtkBox" id="ChatBox"> <property name="name">ChatBox</property> <property name="visible">True</property> <property name="can_focus">False</property> <property name="orientation">vertical</property> <child> <object class="GtkScrolledWindow" id="ScrolledWindow"> <property name="visible">True</property> <property name="can_focus">True</property> <property name="hscrollbar_policy">never</property> <property name="shadow_type">in</property> <child> <object class="GtkViewport" id="viewport1"> <property name="visible">True</property> <property name="can_focus">False</property> <child> <object class="GtkBox" id="Conversation"> <property name="name">Conversation</property> <property name="visible">True</property> <property name="can_focus">False</property> <property name="orientation">vertical</property> <child> <placeholder/> </child> </object> </child> </object> </child> </object> <packing> <property name="expand">True</property> <property name="fill">True</property> <property name="position">0</property> </packing> </child> <child> <object class="GtkBox" id="box2"> <property name="visible">True</property> <property name="can_focus">False</property> <child> <object class="GtkEntry" id="Input"> <property name="visible">True</property> <property name="can_focus">True</property> <property name="has_focus">True</property> </object> <packing> <property name="expand">True</property> <property name="fill">True</property> <property name="position">0</property> </packing> </child> <child> <object class="GtkButton" id="Send"> <property name="label" translatable="yes">Send</property> <property name="visible">True</property> <property name="can_focus">True</property> <property name="receives_default">True</property> </object> <packing> <property name="expand">False</property> <property name="fill">True</property> <property name="position">1</property> </packing> </child> </object> <packing> <property name="expand">False</property> <property name="fill">True</property> <property name="position">1</property> </packing> </child> </object> </child> </object> </interface> `
const ConfigWindow string =` <?xml version="1.0" encoding="UTF-8"?> <!-- Generated with glade 3.18.3 --> <interface> <requires lib="gtk+" version="3.12"/> <object class="GtkWindow" id="Window"> <property name="can_focus">False</property> <property name="title" translatable="yes">Configuration</property> <property name="modal">True</property> <property name="default_width">400</property> <property name="default_height">500</property> <child> <object class="GtkGrid" id="grid1"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="margin_left">5</property> <property name="margin_right">5</property> <property name="margin_top">5</property> <property name="margin_bottom">5</property> <property name="row_spacing">10</property> <property name="column_spacing">10</property> <child> <object class="GtkLabel" id="DataDirLabel"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="label" translatable="yes">Data Directory</property> </object> <packing> <property name="left_attach">0</property> <property name="top_attach">0</property> </packing> </child> <child> <object class="GtkLabel" id="TempDirLabel"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="label" translatable="yes">Temp Directory</property> </object> <packing> <property name="left_attach">0</property> <property name="top_attach">1</property> </packing> </child> <child> <object class="GtkEntry" id="DataDirEntry"> <property name="visible">True</property> <property name="sensitive">False</property> <property name="can_focus">True</property> <property name="hexpand">True</property> </object> <packing> <property name="left_attach">1</property> <property name="top_attach">0</property> <property name="width">3</property> </packing> </child> <child> <object class="GtkEntry" id="TempDirEntry"> <property name="visible">True</property> <property name="sensitive">False</property> <property name="can_focus">True</property> <property name="hexpand">True</property> </object> <packing> <property name="left_attach">1</property> <property name="top_attach">1</property> <property name="width">3</property> </packing> </child> <child> <object class="GtkCheckButton" id="Autologin"> <property name="label" translatable="yes">Autologin</property> <property name="visible">True</property> <property name="can_focus">True</property> <property name="receives_default">False</property> <property name="hexpand">True</property> <property name="xalign">0</property> <property name="draw_indicator">True</property> </object> <packing> <property name="left_attach">0</property> <property name="top_attach">2</property> <property name="width">4</property> </packing> </child> <child> <object class="GtkButton" id="Cancel"> <property name="label" translatable="yes">Cancel</property> <property name="visible">True</property> <property name="can_focus">True</property> <property name="receives_default">True</property> </object> <packing> <property name="left_attach">2</property> <property name="top_attach">3</property> <property name="width">2</property> </packing> </child> <child> <object class="GtkButton" id="Save"> <property name="label" translatable="yes">Save</property> <property name="visible">True</property> <property name="can_focus">True</property> <property name="receives_default">True</property> </object> <packing> <property name="left_attach">0</property> <property name="top_attach">3</property> <property name="width">2</property> </packing> </child> </object> </child> </object> </interface> `
const DownloadableContent string =` <?xml version="1.0" encoding="UTF-8"?> <!-- Generated with glade 3.18.3 --> <interface> <requires lib="gtk+" version="3.12"/> <object class="GtkBox" id="Content"> <property name="visible">True</property> <property name="can_focus">False</property> <child> <object class="GtkEventBox" id="PreviewBox"> <property name="visible">True</property> <property name="can_focus">False</property> <child> <placeholder/> </child> </object> <packing> <property name="expand">False</property> <property name="fill">True</property> <property name="position">0</property> </packing> </child> <child> <object class="GtkMenuButton" id="MenuButton"> <property name="visible">True</property> <property name="can_focus">True</property> <property name="receives_default">True</property> <property name="valign">start</property> <child> <placeholder/> </child> </object> <packing> <property name="expand">False</property> <property name="fill">True</property> <property name="position">1</property> </packing> </child> </object> </interface> `
const DownloadingWindow string =` <?xml version="1.0" encoding="UTF-8"?> <!-- Generated with glade 3.18.3 --> <interface> <requires lib="gtk+" version="3.12"/> <object class="GtkWindow" id="Window"> <property name="can_focus">False</property> <property name="title" translatable="yes">Downloading...</property> <property name="default_width">300</property> <property name="default_height">200</property> <child> <object class="GtkBox" id="box1"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="orientation">vertical</property> <child> <object class="GtkSpinner" id="spinner1"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="active">True</property> </object> <packing> <property name="expand">True</property> <property name="fill">True</property> <property name="position">0</property> </packing> </child> <child> <object class="GtkLabel" id="label1"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="label" translatable="yes">Downloading file</property> </object> <packing> <property name="expand">False</property> <property name="fill">True</property> <property name="position">1</property> </packing> </child> <child> <object class="GtkLabel" id="Filename"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="label" translatable="yes">Filename</property> <property name="wrap">True</property> <property name="wrap_mode">word-char</property> <property name="selectable">True</property> </object> <packing> <property name="expand">False</property> <property name="fill">True</property> <property name="position">2</property> </packing> </child> </object> </child> </object> </interface> `
const FileChooserWindow string =` <?xml version="1.0" encoding="UTF-8"?> <!-- Generated with glade 3.18.3 --> <interface> <requires lib="gtk+" version="3.12"/> <object class="GtkWindow" id="Window"> <property name="can_focus">False</property> <property name="title" translatable="yes">Choose Directory</property> <property name="modal">True</property> <child> <object class="GtkBox" id="box1"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="margin_left">5</property> <property name="margin_right">5</property> <property name="margin_top">5</property> <property name="margin_bottom">5</property> <property name="orientation">vertical</property> <property name="spacing">10</property> <child> <object class="GtkFileChooserWidget" id="FileChooser"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="action">select-folder</property> <property name="do_overwrite_confirmation">True</property> </object> <packing> <property name="expand">True</property> <property name="fill">True</property> <property name="position">0</property> </packing> </child> <child> <object class="GtkBox" id="box2"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="halign">end</property> <property name="spacing">10</property> <child> <object class="GtkButton" id="Cancel"> <property name="label" translatable="yes">Cancel</property> <property name="width_request">100</property> <property name="visible">True</property> <property name="can_focus">True</property> <property name="receives_default">True</property> <property name="halign">center</property> </object> <packing> <property name="expand">False</property> <property name="fill">True</property> <property name="position">0</property> </packing> </child> <child> <object class="GtkButton" id="Save"> <property name="label" translatable="yes">Save</property> <property name="width_request">100</property> <property name="visible">True</property> <property name="can_focus">True</property> <property name="receives_default">True</property> <property name="halign">center</property> </object> <packing> <property name="expand">False</property> <property name="fill">True</property> <property name="position">1</property> </packing> </child> </object> <packing> <property name="expand">False</property> <property name="fill">True</property> <property name="position">1</property> </packing> </child> </object> </child> </object> </interface> `
const FriendsTypeLabel string =` <?xml version="1.0" encoding="UTF-8"?> <!-- Generated with glade 3.18.3 --> <interface> <requires lib="gtk+" version="3.12"/> <object class="GtkLabel" id="FriendsType"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="label" translatable="yes">label</property> <attributes> <attribute name="weight" value="bold"/> <attribute name="scale" value="1.2"/> <attribute name="foreground" value="#34346565a4a4"/> </attributes> </object> </interface> `
const ImageContent string =` <?xml version="1.0" encoding="UTF-8"?> <!-- Generated with glade 3.18.3 --> <interface> <requires lib="gtk+" version="3.12"/> <object class="GtkBox" id="Content"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="orientation">vertical</property> <child> <object class="GtkImage" id="Image"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="stock">gtk-missing-image</property> </object> <packing> <property name="expand">False</property> <property name="fill">True</property> <property name="position">0</property> </packing> </child> </object> </interface> `
const LoginWindow string =` <?xml version="1.0" encoding="UTF-8"?> <!-- Generated with glade 3.18.3 --> <interface> <requires lib="gtk+" version="3.12"/> <object class="GtkWindow" id="Window"> <property name="name">Window</property> <property name="can_focus">False</property> <property name="double_buffered">False</property> <property name="title" translatable="yes">Login</property> <property name="default_width">400</property> <property name="default_height">500</property> <child> <object class="GtkGrid" id="grid1"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="margin_left">5</property> <property name="margin_right">5</property> <property name="margin_top">5</property> <property name="margin_bottom">5</property> <property name="hexpand">True</property> <property name="orientation">vertical</property> <property name="row_spacing">10</property> <child> <object class="GtkLabel" id="IdLabel"> <property name="name">IdLabel</property> <property name="visible">True</property> <property name="can_focus">False</property> <property name="label" translatable="yes">ID</property> </object> <packing> <property name="left_attach">0</property> <property name="top_attach">0</property> </packing> </child> <child> <object class="GtkLabel" id="PasswordLabel"> <property name="name">PasswordLabel</property> <property name="visible">True</property> <property name="can_focus">False</property> <property name="label" translatable="yes">Password</property> </object> <packing> <property name="left_attach">0</property> <property name="top_attach">1</property> </packing> </child> <child> <object class="GtkEntry" id="IdEntry"> <property name="name">IdEntry</property> <property name="visible">True</property> <property name="can_focus">True</property> <property name="is_focus">True</property> <property name="hexpand">True</property> </object> <packing> <property name="left_attach">1</property> <property name="top_attach">0</property> <property name="width">3</property> </packing> </child> <child> <object class="GtkEntry" id="PasswordEntry"> <property name="name">PasswordEntry</property> <property name="visible">True</property> <property name="can_focus">True</property> <property name="hexpand">True</property> <property name="visibility">False</property> <property name="invisible_char">*</property> <property name="input_purpose">alpha</property> </object> <packing> <property name="left_attach">1</property> <property name="top_attach">1</property> <property name="width">3</property> </packing> </child> <child> <object class="GtkButton" id="Login"> <property name="label" translatable="yes">Login</property> <property name="name">LoginButton</property> <property name="visible">True</property> <property name="can_focus">True</property> <property name="receives_default">True</property> <property name="hexpand">True</property> </object> <packing> <property name="left_attach">0</property> <property name="top_attach">3</property> <property name="width">4</property> </packing> </child> <child> <object class="GtkButton" id="Exit"> <property name="label" translatable="yes">Exit</property> <property name="visible">True</property> <property name="can_focus">True</property> <property name="receives_default">True</property> <property name="hexpand">True</property> </object> <packing> <property name="left_attach">0</property> <property name="top_attach">4</property> <property name="width">4</property> </packing> </child> <child> <object class="GtkCheckButton" id="Autologin"> <property name="label" translatable="yes">Autologin</property> <property name="name">Autologin</property> <property name="visible">True</property> <property name="can_focus">True</property> <property name="receives_default">False</property> <property name="xalign">0</property> <property name="draw_indicator">True</property> </object> <packing> <property name="left_attach">0</property> <property name="top_attach">2</property> <property name="width">4</property> </packing> </child> </object> </child> </object> </interface> `
const MainWindow string =` <?xml version="1.0" encoding="UTF-8"?> <!-- Generated with glade 3.18.3 --> <interface> <requires lib="gtk+" version="3.12"/> <object class="GtkWindow" id="Window"> <property name="can_focus">False</property> <property name="title" translatable="yes">Goline</property> <property name="default_width">400</property> <property name="default_height">500</property> <child> <object class="GtkNotebook" id="notebook1"> <property name="visible">True</property> <property name="can_focus">True</property> <child> <object class="GtkScrolledWindow" id="scrolledwindow1"> <property name="visible">True</property> <property name="can_focus">True</property> <property name="shadow_type">in</property> <child> <object class="GtkViewport" id="viewport1"> <property name="visible">True</property> <property name="can_focus">False</property> <child> <object class="GtkBox" id="FriendsBox"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="margin_left">5</property> <property name="margin_right">5</property> <property name="margin_top">5</property> <property name="margin_bottom">5</property> <property name="orientation">vertical</property> <property name="spacing">10</property> <child> <placeholder/> </child> </object> </child> </object> </child> </object> </child> <child type="tab"> <object class="GtkLabel" id="FriendsTab"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="label" translatable="yes">Friends</property> </object> <packing> <property name="tab_fill">False</property> </packing> </child> <child> <object class="GtkBox" id="box2"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="margin_left">5</property> <property name="margin_right">5</property> <property name="margin_top">5</property> <property name="margin_bottom">5</property> <property name="orientation">vertical</property> <property name="spacing">10</property> <child> <object class="GtkButton" id="ConfigButton"> <property name="label" translatable="yes">Configure</property> <property name="visible">True</property> <property name="can_focus">True</property> <property name="receives_default">True</property> </object> <packing> <property name="expand">False</property> <property name="fill">True</property> <property name="position">0</property> </packing> </child> <child> <object class="GtkButton" id="LogoutButton"> <property name="label" translatable="yes">Logout(Revoke)</property> <property name="visible">True</property> <property name="can_focus">True</property> <property name="receives_default">True</property> </object> <packing> <property name="expand">False</property> <property name="fill">True</property> <property name="position">1</property> </packing> </child> <child> <object class="GtkButton" id="ExitButton"> <property name="label" translatable="yes">Exit</property> <property name="visible">True</property> <property name="can_focus">True</property> <property name="receives_default">True</property> </object> <packing> <property name="expand">False</property> <property name="fill">True</property> <property name="position">2</property> </packing> </child> </object> <packing> <property name="position">1</property> </packing> </child> <child type="tab"> <object class="GtkLabel" id="MoreTab"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="label" translatable="yes">More</property> </object> <packing> <property name="position">1</property> <property name="tab_fill">False</property> </packing> </child> <child> <placeholder/> </child> <child type="tab"> <placeholder/> </child> </object> </child> </object> </interface> `
const Message string =` <?xml version="1.0" encoding="UTF-8"?> <!-- Generated with glade 3.18.3 --> <interface> <requires lib="gtk+" version="3.12"/> <object class="GtkBox" id="Message"> <property name="name">Message</property> <property name="visible">True</property> <property name="can_focus">False</property> <property name="halign">start</property> <property name="orientation">vertical</property> <child> <object class="GtkFrame" id="Frame"> <property name="name">MessageFrame</property> <property name="visible">True</property> <property name="can_focus">False</property> <property name="halign">start</property> <property name="label_xalign">0</property> <property name="shadow_type">in</property> <child> <object class="GtkAlignment" id="Alignment"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="left_padding">12</property> <child> <placeholder/> </child> </object> </child> <child type="label"> <object class="GtkLabel" id="Name"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="label" translatable="yes">Name</property> <property name="selectable">True</property> </object> </child> </object> <packing> <property name="expand">False</property> <property name="fill">True</property> <property name="position">0</property> </packing> </child> </object> </interface> `
const PasswordWindow string =` <?xml version="1.0" encoding="UTF-8"?> <!-- Generated with glade 3.18.3 --> <interface> <requires lib="gtk+" version="3.12"/> <object class="GtkWindow" id="Window"> <property name="can_focus">False</property> <property name="modal">True</property> <property name="window_position">center-on-parent</property> <property name="default_width">300</property> <property name="default_height">200</property> <child> <object class="GtkBox" id="box1"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="margin_left">5</property> <property name="margin_right">5</property> <property name="margin_top">5</property> <property name="margin_bottom">5</property> <property name="orientation">vertical</property> <property name="spacing">10</property> <child> <object class="GtkLabel" id="Message"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="label" translatable="yes">Please enter your password</property> </object> <packing> <property name="expand">True</property> <property name="fill">True</property> <property name="position">0</property> </packing> </child> <child> <object class="GtkEntry" id="Password"> <property name="visible">True</property> <property name="can_focus">True</property> <property name="visibility">False</property> <property name="invisible_char">*</property> </object> <packing> <property name="expand">True</property> <property name="fill">True</property> <property name="position">1</property> </packing> </child> <child> <object class="GtkButton" id="Ok"> <property name="label" translatable="yes">OK</property> <property name="width_request">100</property> <property name="visible">True</property> <property name="can_focus">True</property> <property name="receives_default">True</property> <property name="halign">center</property> <property name="valign">center</property> </object> <packing> <property name="expand">True</property> <property name="fill">True</property> <property name="position">2</property> </packing> </child> </object> </child> </object> </interface> `
const TextContent string =` <?xml version="1.0" encoding="UTF-8"?> <!-- Generated with glade 3.18.3 --> <interface> <requires lib="gtk+" version="3.12"/> <object class="GtkBox" id="Content"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="orientation">vertical</property> <child> <object class="GtkLabel" id="Text"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="label" translatable="yes">Text</property> <property name="wrap">True</property> <property name="wrap_mode">word-char</property> <property name="selectable">True</property> </object> <packing> <property name="expand">False</property> <property name="fill">True</property> <property name="position">0</property> </packing> </child> </object> </interface> `
const VerificationWindow string =` <?xml version="1.0" encoding="UTF-8"?> <!-- Generated with glade 3.18.3 --> <interface> <requires lib="gtk+" version="3.12"/> <object class="GtkWindow" id="Window"> <property name="can_focus">False</property> <property name="title" translatable="yes">Verification</property> <property name="default_width">300</property> <property name="default_height">200</property> <child> <object class="GtkBox" id="box1"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="margin_left">5</property> <property name="margin_right">5</property> <property name="margin_top">5</property> <property name="margin_bottom">5</property> <property name="orientation">vertical</property> <property name="spacing">10</property> <child> <object class="GtkLabel" id="Message"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="vexpand">True</property> <property name="label" translatable="yes">Please enter the verification code below into your mobile device.</property> </object> <packing> <property name="expand">False</property> <property name="fill">True</property> <property name="position">0</property> </packing> </child> <child> <object class="GtkLabel" id="Code"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="vexpand">True</property> <property name="label" translatable="yes">code</property> <attributes> <attribute name="weight" value="bold"/> <attribute name="scale" value="1.5"/> <attribute name="foreground" value="#34346565a4a4"/> </attributes> </object> <packing> <property name="expand">False</property> <property name="fill">True</property> <property name="position">1</property> </packing> </child> <child> <object class="GtkButton" id="ExitButton"> <property name="label" translatable="yes">Exit</property> <property name="width_request">100</property> <property name="visible">True</property> <property name="can_focus">True</property> <property name="receives_default">True</property> <property name="halign">center</property> <property name="valign">center</property> <property name="hexpand">False</property> <property name="vexpand">True</property> </object> <packing> <property name="expand">False</property> <property name="fill">True</property> <property name="position">2</property> </packing> </child> </object> </child> </object> </interface> `
const YouMessage string =` <?xml version="1.0" encoding="UTF-8"?> <!-- Generated with glade 3.18.3 --> <interface> <requires lib="gtk+" version="3.12"/> <object class="GtkBox" id="Message"> <property name="name">Message</property> <property name="visible">True</property> <property name="can_focus">False</property> <property name="halign">end</property> <property name="orientation">vertical</property> <child> <object class="GtkFrame" id="Frame"> <property name="name">MessageFrame</property> <property name="visible">True</property> <property name="can_focus">False</property> <property name="halign">start</property> <property name="label_xalign">1</property> <property name="shadow_type">in</property> <child> <object class="GtkAlignment" id="Alignment"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="left_padding">12</property> <child> <placeholder/> </child> </object> </child> <child type="label"> <object class="GtkLabel" id="Name"> <property name="visible">True</property> <property name="can_focus">False</property> <property name="label" translatable="yes">You</property> <property name="selectable">True</property> </object> </child> </object> <packing> <property name="expand">False</property> <property name="fill">True</property> <property name="position">0</property> </packing> </child> </object> </interface> `