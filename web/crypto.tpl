{{ template "head" .}}
<div class="fullScreen crypto">
  
<div class="head">
强大的加解密工具，支持多种算法。
</div>

<div class="input container">
  <div class="row rowtype">
    <div class="col-xs-6 selected" etype="encode">加密</div>
    <div class="col-xs-6" etype="decode">解密</div>
  </div>
  <div class="row rowmethod">
    <div class="col-xs-2 selected">md5</div>
    <div class="col-xs-2">base64</div>
    <div class="col-xs-2">url</div>
    <div class="col-xs-2">sha1</div>
    <div class="col-xs-2">sha256</div>
    <div class="col-xs-2">sha512</div>
  </div>
  <div class="row rowmethod">
    <div class="col-xs-2">aes</div>
    <div class="col-xs-2">des</div>
    <div class="col-xs-2">cipher</div>
    <div class="col-xs-2">dsa</div>
    <div class="col-xs-2">ecdsa</div>
    <div class="col-xs-2">elliptic</div>
  </div>

  <div class="row rowmethod" style="display:none;">
    <div class="col-xs-2">hmac</div>
    <div class="col-xs-2">rand</div>
    <div class="col-xs-2">tls</div>
    <div class="col-xs-2">rc4</div>
    <div class="col-xs-2">rsa</div>
    <div class="col-xs-2">subtle</div>
  </div>

  <input type="text" id="strencode" class="form-control" placeholder="输入要加密或解密的字符串">

  <input type="text" id="strkey" class="form-control" style="display:none;" placeholder="定义一个加密的密钥">
</div>

<div class="output"></div>

</div>
{{template "foot" .}}