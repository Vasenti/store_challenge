# Pregunta por el número del ticket
$ticket_number = Read-Host -Prompt "Introduce el número de ticket (por ejemplo, LOG-123)"

# Pedir mensaje de commit
$commit_message = Read-Host -Prompt "Ingrese el mensaje de commit"

# Pregunta por el tipo de commit
Write-Host "Selecciona el tipo de commit (1 = feature, 2 = hotfix, 3 = bug):"
Write-Host "1) Feature (✨)"
Write-Host "2) Hotfix (🚑)"
Write-Host "3) Bug (🐛)"

$commit_type = Read-Host -Prompt "Opción"

# Según la opción seleccionada, añadimos el emoji correspondiente
switch ($commit_type) {
    1 { $commit_type_text = "✨" }
    2 { $commit_type_text = "🚑" }
    3 { $commit_type_text = "🐛" }
    default {
        Write-Host "Opción inválida. El commit será tratado como una 'Feature'."
        $commit_type_text = "✨"
    }
}

# Concatenar el mensaje del commit
$final_message = "$commit_type_text [$ticket_number] - $commit_message"

# Mostrar el mensaje final
Write-Host "Commit ejecutado con el siguiente mensaje: $final_message"

# Ejecutar el commit
git commit -m "$final_message"