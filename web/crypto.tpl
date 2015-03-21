{{ template "head" .}}
<div class="fullScreen crypto">
<div class="head">
强大的加解密工具，支持多种算法。
</div>
<div class="input">
  <div class="row">
    <div class="col-xs-6 selected">加密</div>
    <div class="col-xs-6">解密</div>
  </div>
  <div class="row">
    <div class="col-xs-2 selected">md5</div>
    <div class="col-xs-2">base64</div>
    <div class="col-xs-2">url</div>
    <div class="col-xs-2">sha1</div>
    <div class="col-xs-2">sha256</div>
    <div class="col-xs-2">sha512</div>
  </div>
  <div class="row">
    <div class="col-xs-2">aes</div>
    <div class="col-xs-2">des</div>
    <div class="col-xs-2">cipher</div>
    <div class="col-xs-2">dsa</div>
    <div class="col-xs-2">ecdsa</div>
    <div class="col-xs-2">elliptic</div>
  </div>

  <div class="row">
    <div class="col-xs-2 selected">hmac</div>
    <div class="col-xs-2">rand</div>
    <div class="col-xs-2">tls</div>
    <div class="col-xs-2">rc4</div>
    <div class="col-xs-2">rsa</div>
    <div class="col-xs-2">subtle</div>
  </div>
  
  <input type="text" id="strencode" class="form-control">
</div>

<div class="result">

</div>
</div>
{{template "foot" .}}