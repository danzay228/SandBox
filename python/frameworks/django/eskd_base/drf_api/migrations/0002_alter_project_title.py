# Generated by Django 4.1.1 on 2022-10-17 11:47

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('drf_api', '0001_initial'),
    ]

    operations = [
        migrations.AlterField(
            model_name='project',
            name='title',
            field=models.CharField(max_length=40),
        ),
    ]
