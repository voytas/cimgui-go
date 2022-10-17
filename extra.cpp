#include "extra.h"

extern "C" const char *getClipboardText(void* user_data);
extern "C" void setClipboardText(void* user_data, const char *text);

void RegisterClipboardFunctions(ImGuiIO *self) {
   ImGuiIO &io = *reinterpret_cast<ImGuiIO *>(self);
   io.GetClipboardTextFn = getClipboardText;
   io.SetClipboardTextFn = setClipboardText;
}
