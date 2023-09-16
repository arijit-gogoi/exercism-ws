;;; word-count.el --- word-count Exercise (exercism)  -*- lexical-binding: t; -*-

;;; Commentary:

;;; Code:

(defun word-count (str)
  (let ((words (split-string
                (downcase
                 (replace-regexp-in-string " '+\\|'+ \\|^'+\\|'+$" ""
                                           (replace-regexp-in-string "[^[:alnum:]' ]+" " " str)))))
        result)
    (dolist (word words)
      (cl-incf (alist-get word result 0 nil 'string=)))
    result))

(provide 'word-count)
;;; word-count.el ends here
