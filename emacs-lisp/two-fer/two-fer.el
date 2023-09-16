;;; two-fer.el --- Two-fer Exercise (exercism)  -*- lexical-binding: t; -*-

;;; Commentary:

;;; Code:
(defun two-fer (&optional name)
  (format "One for %s, one for me." (or name "you")))

;; this also works
;; (cl-defun two-fer (&optional (name "you"))
;;     (format "One for %s, one for me." name))

(provide 'two-fer)
;;; two-fer.el ends here
