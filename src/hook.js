//获取WeChatAppEx.exe的基址
var module = Process.findModuleByName("WeChatAppEx Framework");
var base = module.base;

// var address = {
//   enableVconsole: "0x25505f8",
//   setupInterceptor: "0x23f028c",
//   adevTools: "0x6c219f8",
//   version: "WeChat v3.8.8.18",
// };

var address = {
  enableVconsole: "0x2550414",
  setupInterceptor: "0x23f00a8",
  adevTools: "0x6c218f8",
  version: "WeChat v3.8.9",
};

send("WeChatAppEx 注入成功!");
send("当前小程序版本: " + address.version);
send("等待小程序加载...");

/* 公共部分 */
function readStdString(s) {
  var flag = s.add(23).readU8();
  if (flag == 0x80) {
    // 从堆中读取
    var size = s.add(8).readUInt();
    return s.readPointer().readUtf8String(size);
  } else {
    // 从栈中读取
    return s.readUtf8String(flag);
  }
}
function writeStdString(s, content) {
  var flag = s.add(23).readU8();
  if (flag == 0x80) {
    // 从堆中写入
    var orisize = s.add(8).readUInt();
    if (content.length > orisize) {
      throw "must below orisize!";
    }
    s.readPointer().writeUtf8String(content);
    s.add(8).writeUInt(content.length);
  } else {
    // 从栈中写入
    if (content.length > 22) {
      throw "max 23 for stack str";
    }
    s.writeUtf8String(content);
    s.add(23).writeU8(content.length);
  }
}

/** hook 部分 **/

// 开启 enable_vconsole
function enableVconsole() {
  Interceptor.attach(base.add(address.enableVconsole), {
    onEnter(args) {
      for (var i = 0; i < 0x1000; i += 8) {
        try {
          var s = readStdString(args[2].add(i));
          var s1 = s.replaceAll(
            '"enable_vconsole":false',
            '"enable_vconsole": true'
          );
          if (s !== s1) {
            writeStdString(args[2].add(i), s1);
          }
        } catch (a) {
          // console.log(a);
        }
      }
    },
  });
}

//  开启 devtools 所有功能
function setupInterceptor() {
  Interceptor.attach(base.add(address.setupInterceptor), {
    onEnter(args) {
      args[1] = ptr(0x1);
      send("已还原完整F12");
    },
  });
}

// 解除 仅首次启动微信打开 devTools

function bypassOnlyFirst() {
  var menuItemDevTools = base.add(address.adevTools);
  Memory.protect(menuItemDevTools, 8, "rw-");
  menuItemDevTools.writeUtf8String("DevTools");
  enableVconsole();
  setupInterceptor();
}

bypassOnlyFirst();
