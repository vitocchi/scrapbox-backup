# scrapbox-backup
KMSでGitHubTokenを暗号化しているので、暗号化したTOKENをそのままgithubに上げれちゃう。便利。

1. GitHubでAccess Tokenを取得

2. aws cliでkmsのAPIを叩いてAccess Tokenを暗号化する

3. ソースコード内でkmsのSDKを使って暗号化したTOKENを復号化する

4. Lambda関数にKMSを利用するロールをもたせることでLambdaは復号化APIを叩くことができる

5. KMSを利用するロールを持っていなければAccess Tokenを復号化出来ない。
