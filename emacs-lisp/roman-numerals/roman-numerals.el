;;; roman-numerals.el --- roman-numerals Exercise (exercism)  -*- lexical-binding: t; -*-

;;; Commentary:
;;; Convert from normal numbers to Roman Numerals.


;;; Code:
(defun to-roman-base (number &optional unit five decade thousand)
  "Convert NUMBER using UNIT FIVE and DECADE and THOUSAND."
  (cond
   ((zerop number) "")
   ((= number 4) (string unit five))
   ((= number 9) (string unit decade))
   ((< number 4) (make-string number unit))
   ((< number 9) (concat (string five)
                         (make-string (- number 5) unit)))
   (t (make-string number thousand))))

(defun to-roman (number)
  "Convert NUMBER to roman."
  (concat
   (to-roman-base (/ (mod number 10000) 1000) ?M)
   (to-roman-base (/ (mod number 1000) 100) ?C ?D ?M)
   (to-roman-base (/ (mod number 100) 10) ?X ?L ?C)
   (to-roman-base (mod number 10 ) ?I ?V ?X)))


(provide 'roman-numerals)
;;; roman-numerals.el ends here
