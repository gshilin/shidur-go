class window.BigWindow
  bigWindow: window.open('/big_windows', 'Big Window', 'height="' + screen.height + '",width="' + screen.width + '",titlebar=no,fullscreen=yes,menubar=no,location=no,resizable=yes,scrollbars=no,status=no')

  show_slide: true

  displayLiveSlide: (content) =>
    $(@bigWindow.document.body).find(".content").html(content)

  displayLiveQuestion: (content) =>
    $(@bigWindow.document.body).find(".question").html(content)

  switchSlidesQuestion: =>
    question = $(@bigWindow.document.body).find(".question")[0]
    slide = $(@bigWindow.document.body).find(".slides")[0]
    @show_slide = !@show_slide;
    slide.style.display = if this.show_slide then 'block' else 'none'
    question.style.display = if this.show_slide then 'none' else 'block'
