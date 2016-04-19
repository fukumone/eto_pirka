$(function(){
  $(".confirmation").click(function(){
    if (!confirm('データを削除します\nよろしいですか？')) {
      return false
    }
  })
})
