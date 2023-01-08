
"Vim Base
" 行号
set number
" 高亮
syntax enable
" 编码
set fileencodings=utf-8,ucs-bom,gb18030,gbk,gb2312,cp936
set termencoding=utf-8
set encoding=utf-8

" 自动缩进
set autoindent
" Tab键空格位
set tabstop=4
" 突出显示当前行
set cursorline
" 开启实时搜索
set incsearch
" 搜索时大小写不敏感
set ignorecase
" 插件自动补全
filetype plugin indent on 
" 退出插入模式指定类型的文件自动保存
au InsertLeave *.go,*.sh,*.php write

call plug#begin('~/.config/nvim/plugged')

" 主题
"Plug 'pineapplegiant/spaceduck', { 'branch': 'main' }
"
" 在Plugged目录下手动克隆比较快 安装 mv .vim/autoload/plugged/vim-hybrid/colors/  .vim    mv ~/.vim/colors ~/.config/nvim
Plug 'chriskempson/base16-vim'
" 语法高亮
Plug 'sheerun/vim-polyglot'
" TAB
"Plug 'ervandew/supertab'
" 括号
Plug 'jiangmiao/auto-pairs'
" 底部状态栏美化
Plug 'vim-airline/vim-airline'
Plug 'vim-airline/vim-airline-themes'
" 彩虹括号
Plug 'luochen1990/rainbow'
" Go 需要 :GoInstallBinaries  手动安装看下面
Plug 'fatih/vim-go', { 'do': ':GoUpdateBinaries' }

" 代码提示 需要nodejs sudo apt install nodejs npm
"Plug 'neoclide/coc.nvim', {'branch': 'release'}

" 侧边栏
Plug 'preservim/nerdtree'

" 代码自动完成，安装完插件还需要额外配置才可以使用
Plug 'Valloric/YouCompleteMe'

" go 中的代码追踪，输入 gd 就可以自动跳转
Plug 'dgryski/vim-godef'

call plug#end()

" 主题配置 =============================
" 第一套黑色
"if exists('+termguicolors')
"   let &t_8f = "\<Esc>[38;2;%lu;%lu;%lum"
"   let &t_8b = "\<Esc>[48;2;%lu;%lu;%lum"
"  set termguicolors
"endif
" colorscheme spaceduck
" 第二套
"Plug 'chriskempson/base16-vim'
set background=dark
colorscheme hybrid


" 语法高亮===========================
set nocompatible

" 底栏主题 =============================

" 开启24bit的颜色，开启这个颜色会更漂亮一些
set termguicolors
" 命令行开启
"" 配色方案, 可以从上面插件安装中的选择一个使用
set background=dark " 主题背景 dark-深色; light-浅u
" 官网 https://github.com/vim-airline/vim-airline-themes
" :AirlineTheme luna


"GO  ===================================
" 注意修改代理 go get -v golang.org/x/tools/gopls    不行的话手动下载
" https://www.sunzhongwei.com/amp/vim-execution-goinstallbinaries-installation-depend-on-failure
let g:go_def_mode='gopls'
let g:go_info_mode='gopls'



"==============================================================================
"  Valloric/YouCompleteMe 插件
"==============================================================================
let g:ycm_key_list_select_completion = ['<C-n>', '<space>']
let g:ycm_key_list_previous_completion = ['<C-p>', '<Up>']
let g:SuperTabDefaultCompletionType = '<C-n>'

" better key bindings for UltiSnipsExpandTrigger
let g:UltiSnipsExpandTrigger = "<tab>"
let g:UltiSnipsJumpForwardTrigger = "<tab>"
let g:UltiSnipsJumpBackwardTrigger = "<s-tab>"

